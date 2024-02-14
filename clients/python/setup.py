"""Setup script for Multiparty package."""
from setuptools import setup, find_packages

with open('README.md', 'r', encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='multiparty',
    version='0.1.0',
    description='Fastest way to implement approvals and human prompts in your app.',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/indentinc/multiparty',
    packages=find_packages(),
    install_requires=[
        'googleapis-common-protos',
        'langchain',
        'protobuf',
    ],
    classifiers=[
        'License :: OSI Approved :: Apache Software License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
    ],
    python_requires='>=3.9',
)
